{ pkgs ? import <nixpkgs> { system = "x86_64-linux"; } }: with pkgs;
let moby-dick = import ./default.nix { };
in
dockerTools.buildImage {
  name = "moby-dick";
  tag = "9a213bec1b81c0ed1a0bc7b854be4e678ae024d6";
  contents = [ moby-dick cacert ];
  
  extraCommands = ''
    mkdir -p db
  '';

  config = {
    Cmd = [ "${moby-dick}/bin/moby-dick" ];
    Env = [
      "SSL_CERT_FILE=${cacert}/etc/ssl/certs/ca-bundle.crt"
      "DB_PATH=db"
    ];
  };
}