{
  description = "My projects, notes, and labs for CS113";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";

    utils = {
      url = "github:numtide/flake-utils";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs = { self, nixpkgs, utils }:
    utils.lib.eachDefaultSystem
      (system:
        let pkgs = import nixpkgs {
          inherit system;
        };
        in
        {
          devShell = pkgs.mkShell rec {
            name = "cs113-stuffs";

            packages = with pkgs; [
              # Notes
              python39Packages.watchdog
              pandoc
              texlive.combined.scheme-full
            ];
          };
        });
}
