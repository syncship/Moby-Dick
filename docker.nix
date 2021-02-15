{ pkgs ? import <nixpkgs> { system = "x86_64-linux"; } }: with pkgs;
let moby-dick = import ./default.nix { };
in
dockerTools.buildImage {
  name = "moby-dick";
  tag = "fdf0f1afd107c6a19804241c87ade18b01619615";
  contents = [ moby-dick ];
  
  config = {
    Cmd = [ "${moby-dick}/bin/moby-dick" ];
  };
}