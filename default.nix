{ pkgs ? import <nixpkgs> { } }: with pkgs;

buildGoModule rec {
  pname = "moby-dick";
  version = "40c3e4f70ac7b50f35354b1c96ce345fa950d7da";

  src = fetchFromGitHub {
    owner = "syncship";
    repo = "moby-dick";
    rev = "${version}";
    sha256 = "1fphd9xmc8j7n7vg3f0lcvpfr7fqvhng5va8pn7vw7k7ys4j64l1";
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
