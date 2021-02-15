{ pkgs ? import <nixpkgs> { } }: with pkgs;

buildGoModule rec {
  pname = "moby-dick";
  version = "3881b862e73323df5e8303e0d1a42e15ce0bae41";

  src = fetchFromGitHub {
    owner = "syncship";
    repo = "moby-dick";
    rev = "${version}";
    sha256 = "1fphd9xmc8j7n7vg3f0lcvpfr7fqvhng5va8pn7vw7k7ys4j64l2";
  };

  vendorSha256 = "0dxi1358bqjsfc8mivmkwvz35pdzb5icfkyczn7ylgdz0m9qc8p0";
  modSha256 = "1879j77k96684wi554rkjxydrj8g3hpp0kvxz03sd8dmwr3lh83j";

  meta = with lib; {
    description = "a bot with good literary reference";
    homepage = https://github.com/syncship/moby-dick;
    license = licenses.agpl3;
    platforms = platforms.linux ++ platforms.darwin;
  };
}
