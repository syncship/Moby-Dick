{ pkgs ? import <nixpkgs> { } }: with pkgs;

buildGoModule rec {
  pname = "moby-dick";
  version = "9a213bec1b81c0ed1a0bc7b854be4e678ae024d6";

  src = fetchFromGitHub {
    owner = "syncship";
    repo = "moby-dick";
    rev = "${version}";
    sha256 = "0c7q11jjadmiv5g776xlbqi709vpbcm21lzk8ycalay80shppr7d";
  };

  vendorSha256 = "0dxi1358bqjsfc8mivmkwvz35pdzb5icfkyczn7ylgdz0m9qc8p0";
  modSha256 = "1879j77k96684wi554rkjxydrj8g3hpp0kvxz03sd8dmwr3lh83j";

  Cmd = [ "moby-dick" ];

  meta = with lib; {
    description = "a bot with good literary reference";
    homepage = https://github.com/syncship/moby-dick;
    license = licenses.agpl3;
    platforms = platforms.linux ++ platforms.darwin;
  };
}
