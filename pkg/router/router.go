package router

import (
	"errors"
	"fmt"
	"log"
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
type Arguments map[string]typeguard.ArgumentConstructor

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
		return cmd, errors.New(executioner.ErrCommandNotFound)
	}

	// Checks user permission
	p, err := s.State.UserChannelPermissions(m.Author.ID, m.ChannelID)

	if err != nil || (p&int64(helper.Sum(cmd.Permissions)) !=
		int64(helper.Sum(cmd.Permissions))) {
		return cmd, errors.New(executioner.ErrNoPermission)
	}

	// Resets command args
	tempArgs := []string{}

	for _, s := range helper.Reverse(ctx) {
		if strings.HasPrefix(s, r.argPrefix) {
			argName := strings.TrimPrefix(s, r.argPrefix)

			if arg, ok := cmd.Args[argName]; ok {
				switch arg.To {
				case typeguard.WantInt():
					if len(tempArgs) > 1 {
						return cmd, fmt.Errorf(executioner.ErrTooManyValues, argName)
					}

					arg.Output.Value = tempArgs[0]

				case typeguard.WantArrInt():
					arg.Output.Value = strings.Join(tempArgs, ",")

				case typeguard.WantArrString():
					arg.Output.Value = strings.Join(tempArgs, ",")

				default:
					if len(tempArgs) > 1 {
						return cmd, fmt.Errorf(executioner.ErrTooManyValues, argName)
					}

					arg.Output.Value = tempArgs[0]
				}

				cmd.Args[argName] = arg
			}
		}

		tempArgs = append(tempArgs, s)
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
}
