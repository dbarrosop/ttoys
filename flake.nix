{

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";
    nix-filter.url = "github:numtide/nix-filter";
    flake-utils.url = "github:numtide/flake-utils";
    flake-compat = {
      url = "github:edolstra/flake-compat";
      flake = false;
    };
  };

  outputs = { self, nixpkgs, flake-utils, nix-filter, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        overlays = [
          (import ./build/nix/overlay.nix)
        ];

        pkgs = import nixpkgs {
          inherit system overlays;
        };

        src = nix-filter.lib.filter {
          root = ./.;
        };

        nix-src = nix-filter.lib.filter {
          root = ./.;
          include = [
            (nix-filter.lib.matchExt "nix")
          ];
        };

        checkInputs = with pkgs; [
          richgo
          gofumpt
          golangci-lint
          golines
          govulncheck
        ];

        buildInputs = with pkgs; [
        ];

        nativeBuildInputs = with pkgs; [
          go
          clang
        ];

        tags = [ ];

        name = "ttoys";
        description = "Dev Aid Tool";
        version = nixpkgs.lib.fileContents ./VERSION;
        module = "github.com/dbarrosop/ttoys";

        ldflags = [
          "-X main.Version=${version}"
        ];

      in
      {
        checks = {
          nixpkgs-fmt = pkgs.runCommand "check-nixpkgs-fmt"
            {
              nativeBuildInputs = with pkgs;
                [
                  nixpkgs-fmt
                ];
            }
            ''
              mkdir $out
              nixpkgs-fmt --check ${nix-src}
            '';

          ttoys = pkgs.runCommand "gotests"
            {
              nativeBuildInputs = with pkgs; [
              ] ++ checkInputs ++ buildInputs ++ nativeBuildInputs;
            }
            ''
              export GOLANGCI_LINT_CACHE=$TMPDIR/.cache/golangci-lint
              export GOCACHE=$TMPDIR/.cache/go-build
              export GOMODCACHE="$TMPDIR/.cache/mod"
              export GOPATH="$TMPDIR/.cache/gopath"

              echo "➜ Source: ${src}"
              cd ${src}

              echo "➜ Running code formatters, if there are changes, fail"
              golines -l --base-formatter=gofumpt . | diff - /dev/null

              echo "➜ Checking for vulnerabilities"
              govulncheck ./...

              echo "➜ Running golangci-lint"
              golangci-lint run \
                --timeout 300s \
                ./...

              echo "➜ Running tests"
              richgo test \
                -tags="${pkgs.lib.strings.concatStringsSep " " tags}" \
                -ldflags="${pkgs.lib.strings.concatStringsSep " " ldflags}" \
                -v ./...

              mkdir $out

            '';

        };

        devShells = flake-utils.lib.flattenTree rec {
          default = pkgs.mkShell {
            buildInputs = with pkgs; [
            ] ++ checkInputs ++ buildInputs ++ nativeBuildInputs;
          };
        };

        packages = flake-utils.lib.flattenTree rec {
          ttoys = pkgs.buildGoModule {
            inherit name src version ldflags buildInputs nativeBuildInputs;

            vendorSha256 = null;

            doCheck = false;

            meta = with pkgs.lib; {
              description = description;
              homepage = "https://github.com/dbarrosop/ttoys";
              maintainers = [ "dbarrosop" ];
              platforms = platforms.linux ++ platforms.darwin;
            };

          };

          default = ttoys;

        };

      }



    );


}
