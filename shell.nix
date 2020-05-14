with import <nixpkgs> {};

mkShell {
  buildInputs = [
    go_1_12
    plantuml
  ];

  shellHook = ''
    export GOPATH=$HOME/go
  '';
}
