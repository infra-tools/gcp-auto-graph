with import <nixpkgs> {};

buildGo112Module {
  name = "gcp-auto-graph";

  src = ./.;

  modSha256 = null;
}
