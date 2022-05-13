{
  description = "virtual environments";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";

    devshell.url = "github:numtide/devshell";
    devshell.inputs.nixpkgs.follows = "nixpkgs";

    utils.url = "github:numtide/flake-utils";
    utils.inputs.nixpkgs.follows = "nixpkgs";
  };
  
  outputs = { self, utils, devshell, nixpkgs }:
    utils.lib.eachDefaultSystem (system: 
      let
        pkgs = import nixpkgs {
          inherit system;
          config.allowUnfree = true;
          overlays = [ devshell.overlay ];
        };

        drum-hero = pkgs.buildGoModule {
          pname = "drum-hero";
          version = "0.0.1";
          src = ./client;

          vendorSha256 = "sha256-BVPK8jSy9TnvlZm0lVJ09fYS57U48fPU27ur3x3LYVU=";
        };
      in rec {
        packages = utils.lib.flattenTree {
          inherit drum-hero;
        };
        defaultPackage = packages.drum-hero;
        apps.drum-hero = utils.lib.mkApp { drv = packages.drum-hero; };
        defaultApp = apps.drum-hero;
        devShell = pkgs.devshell.mkShell {
          imports = [ (pkgs.devshell.importTOML ./devshell.toml) ];
        };
      }
    );
}
