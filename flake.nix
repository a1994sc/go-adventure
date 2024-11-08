{
  description = "A basic gomod2nix flake";

  inputs = {
    # keep-sorted start block=yes case=no
    flake-parts = {
      url = "github:hercules-ci/flake-parts";
      inputs.nixpkgs-lib.follows = "nixpkgs";
    };
    flake-utils.url = "github:numtide/flake-utils";
    gomod2nix = {
      url = "github:nix-community/gomod2nix";
      inputs.nixpkgs.follows = "nixpkgs";
      inputs.flake-utils.follows = "flake-utils";
    };
    nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";
    pre-commit-hooks = {
      inputs.nixpkgs.follows = "nixpkgs";
      inputs.nixpkgs-stable.follows = "nixpkgs";
      url = "github:/cachix/git-hooks.nix";
    };
    systems.url = "github:nix-systems/default";
    treefmt-nix = {
      url = "github:numtide/treefmt-nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
    # keep-sorted end
  };

  outputs =
    inputs@{
      self,
      nixpkgs,
      ...
    }:
    inputs.flake-parts.lib.mkFlake { inherit inputs; } {
      systems = import inputs.systems;
      perSystem =
        {
          self',
          pkgs,
          lib,
          system,
          ...
        }:
        let
          pkgs = import nixpkgs {
            inherit system;
            overlays = [ inputs.gomod2nix.overlays.default ];
          };
          treefmtEval = inputs.treefmt-nix.lib.evalModule pkgs (
            { pkgs, ... }:
            {
              projectRootFile = "flake.nix";
              # keep-sorted start block=yes case=no
              programs.dprint = {
                enable = true;
                settings = {
                  includes = [
                    "**/*.toml"
                    "**/*.json"
                    "**/*.md"
                  ];
                  excludes = [
                    "**/target"
                  ];
                  plugins =
                    let
                      dprintWasmPluginUrl = n: v: "https://plugins.dprint.dev/${n}-${v}.wasm";
                    in
                    [
                      (dprintWasmPluginUrl "json" "0.19.3")
                      (dprintWasmPluginUrl "markdown" "0.17.8")
                      (dprintWasmPluginUrl "toml" "0.6.2")
                    ];
                };
              };
              programs.jsonfmt = {
                enable = true;
                package = pkgs.jsonfmt;
              };
              programs.keep-sorted.enable = true;
              programs.nixfmt = {
                enable = true;
                package = pkgs.nixfmt-rfc-style;
              };
              programs.statix.enable = true;
              # keep-sorted end
              settings.formatter = {
                # keep-sorted start block=yes
                actionlint = {
                  command = pkgs.actionlint;
                  includes = [ "./.github/workflows/*.yml" ];
                };
                jsonfmt.includes = [
                  "*.json"
                  "./.github/*.json"
                  "./.vscode/*.json"
                ];
                # keep-sorted end
              };
            }
          );
          goEnv = pkgs.mkGoEnv { pwd = ./.; };
        in
        rec {
          devShells.default = pkgs.mkShell {
            shellHook = ''
              GOROOT="$(dirname $(dirname $(which go)))/share/go"
              unset GOPATH;
            '';
            packages = [
              goEnv
              pkgs.go
              pkgs.gopls
              pkgs.gotools
              pkgs.go-tools
              packages.gomod2nix
            ];
          };
          packages.default = pkgs.buildGoApplication {
            pname = "myapp";
            version = "0.1";
            pwd = ./.;
            src = ./.;
            modules = ./gomod2nix.toml;

            ldflags = [
              "-s"
              "-w"
            ];
          };
          packages.gomod2nix = inputs.gomod2nix.packages.${system}.default.overrideAttrs (
            finalAttrs: previousAttrs: {
              patches = [ ./patches/gomod2nix-fix.patch ];
            }
          );
          formatter = treefmtEval.config.build.wrapper;
        };
    };
}
