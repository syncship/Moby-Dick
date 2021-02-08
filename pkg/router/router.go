package router

import (
	"log"
	"sort"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/syncship/moby-dick/pkg/executioner"
	"github.com/syncship/moby-dick/pkg/helper"
	"github.com/syncship/moby-dick/pkg/typeguard"
)

// Router defines the structure of the Router
type Router struct {
	Commands  map[string]Command
	prefix    string
	argPrefix string
}

// Command defines the structure of a command
type Command struct {
	Name        string
	Args        Arguments
	Run         Callback
	Permissions []int
}

// Arguments ..
type Arguments map[string]ArgumentConstructor

// ArgumentConstructor ..
type ArgumentConstructor struct {
	To       string
	Required bool
	Output   typeguard.Output
}

// Callback ..
type Callback func(s *discordgo.Session, m *discordgo.MessageCreate, a Arguments)

// New returns a new Router
func New() *Router {
	return &Router{
		Commands:  map[string]Command{},
		prefix:    ";",
		argPrefix: "--",
	}
}

// SetPrefix sets the command prefix
func (r *Router) SetPrefix(prefix string) {
	r.prefix = prefix
}

// AddCommand Adds a new command to the router
func (r *Router) AddCommand(c Command) {
	if _, ok := r.Commands[c.Name]; ok {
		log.Fatalf("command '%s' has been declared already\n", c.Name)
	}

	r.Commands[c.Name] = c
}

func (r *Router) parseCommand(
	s *discordgo.Session,
	m *discordgo.MessageCreate) (cmd Command, err error) {
	ctx := strings.Split(m.Content, " ")
	name, ctx := strings.TrimPrefix(ctx[0], r.prefix), ctx[1:]

	cmd, ok := r.Commands[name]
	if !ok {
		return cmd, executioner.ErrCommandNotFound{Command: name}
	}

	// Checks user permission
	p, err := s.State.UserChannelPermissions(m.Author.ID, m.ChannelID)

	if err != nil || (p&int64(helper.Sum(cmd.Permissions)) !=
		int64(helper.Sum(cmd.Permissions))) {
		return cmd, executioner.ErrNoPermission{}
	}

	// Stores which args are required for given command
	requiredArgs := []string{}
	for k := range cmd.Args {
		if cmd.Args[k].Required {
			requiredArgs = append(requiredArgs, k)
		}
	}

	// Resets command args
	tempArgs := []string{}

	for _, s := range helper.Reverse(ctx) {
		if !strings.HasPrefix(s, r.argPrefix) {
			tempArgs = append(tempArgs, s)
			continue
		}
		argName := strings.TrimPrefix(s, r.argPrefix)

		foundRequiredArg := sort.SearchStrings(requiredArgs, argName)
		if foundRequiredArg < len(requiredArgs) &&
			requiredArgs[foundRequiredArg] == argName {
			requiredArgs = helper.RemoveString(requiredArgs, argName)
		}

		arg, ok := cmd.Args[argName]
		if !ok {
			tempArgs = append(tempArgs, s)
			continue
		}
		switch arg.To {
		case typeguard.WantInt():
			if len(tempArgs) != 1 {
				return cmd, executioner.ErrValuesOutOfBounds{ArgName: argName}
			}

			arg.Output.Value = tempArgs[0]

		case typeguard.WantArrInt():
			if len(tempArgs) == 0 {
				return cmd, executioner.ErrValuesOutOfBounds{ArgName: argName}
			}

			arg.Output.Value = strings.Join(tempArgs, ",")

		case typeguard.WantArrString():
			if len(tempArgs) == 0 {
				return cmd, executioner.ErrValuesOutOfBounds{ArgName: argName}
			}

			arg.Output.Value = strings.Join(tempArgs, ",")

		default:
			if len(tempArgs) != 1 {
				return cmd, executioner.ErrValuesOutOfBounds{ArgName: argName}
			}

			arg.Output.Value = tempArgs[0]
		}

		cmd.Args[argName] = arg

		tempArgs = append(tempArgs, s)
	}

	if len(requiredArgs) > 0 {
		return cmd, executioner.ErrMissingRequiredArgs{Args: requiredArgs}
	}

	return cmd, nil
}

// OnMessageCreateHandler handles the message create event
func (r *Router) OnMessageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID ||
		!strings.HasPrefix(m.Content, r.prefix) {
		return
	}

	cmd, err := r.parseCommand(s, m)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}

	cmd.Run(s, m, cmd.Args)

	// Resets cmd arguments
	for k := range cmd.Args {
		arg := cmd.Args[k]
		arg.Output.Value = ""

		cmd.Args[k] = arg
	}
}
