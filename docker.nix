{ pkgs ? import <nixpkgs> { system = "x86_64-linux"; } }: with pkgs;
let moby-dick = import ./default.nix { };
in
dockerTools.buildImage {
  name = "moby-dick";
  tag = "3881b862e73323df5e8303e0d1a42e15ce0bae41";
  contents = [ moby-dick ];
  
  config = {
    Cmd = [ "${moby-dick}/bin/moby-dick" ];
  };
}