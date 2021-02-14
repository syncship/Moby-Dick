{ pkgs ? import <nixpkgs> { system = "x86_64-linux"; } }: with pkgs;
let moby-dick = import ./default.nix { };
in
dockerTools.buildLayeredImage {
  name = "moby-dick";
  tag = "40c3e4f70ac7b50f35354b1c96ce345fa950d7da";
  contents = [ moby-dick ];
}
