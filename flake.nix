{
  description = "A very basic flake";

  inputs.nixpkgs.url = "github:nixos/nixpkgs/a1007637cea374bd1bafd754cfd5388894c49129";

  outputs = { self, nixpkgs }: let
    pkgs = nixpkgs.legacyPackages.x86_64-linux;
  in {

    defaultPackage.x86_64-linux = pkgs.buildGoModule {
      name = "gcp-auto-graph";

      src = ./.;

      buildInputs = [ pkgs.makeWrapper ];

      subPackages = [ "./cmd/gcp-auto-graph/main.go" ];

      vendorSha256 = "5gUmLtXogQgb+K7rQLxXNUlj73EstBiqVPHyygBZDg8=";

      postInstall = ''
        mv $out/bin/main $out/bin/gcp-auto-graph
        wrapProgram $out/bin/gcp-auto-graph --prefix PATH : ${pkgs.lib.makeBinPath [ pkgs.plantuml ]}
      '';
    };

    devShell.x86_64-linux = pkgs.mkShell {
      buildInputs = [
        pkgs.go
        pkgs.plantuml
        pkgs.gopls
      ];
    };
  };
}
