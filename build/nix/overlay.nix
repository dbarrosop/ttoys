final: prev: rec {
  golangci-lint = prev.golangci-lint.override rec {
    buildGoModule = args: final.buildGoModule (args // rec {
      version = "1.50.1";
      src = final.fetchFromGitHub {
        owner = "golangci";
        repo = "golangci-lint";
        rev = "v${version}";
        sha256 = "sha256-7HoneQtKxjQVvaTdkjPeu+vJWVOZG3AOiRD87/Ntgn8=";
      };
      vendorSha256 = "sha256-6ttRd2E8Zsf/2StNYt6JSC64A57QIv6EbwAwJfhTDaY=";

      meta = with final.lib; args.meta // {
        broken = false;
      };
    });
  };

  golines = final.buildGoModule rec {
    name = "golines";
    version = "0.11.0";
    src = final.fetchFromGitHub {
      owner = "dbarrosop";
      repo = "golines";
      rev = "b7e767e781863a30bc5a74610a46cc29485fb9cb";
      sha256 = "sha256-pxFgPT6J0vxuWAWXZtuR06H9GoGuXTyg7ue+LFsRzOk=";
    };
    vendorSha256 = "sha256-rxYuzn4ezAxaeDhxd8qdOzt+CKYIh03A9zKNdzILq18=";

    meta = with final.lib; {
      description = "A golang formatter that fixes long lines";
      homepage = "https://github.com/segmentio/golines";
      maintainers = [ "nhost" ];
      platforms = platforms.linux ++ platforms.darwin;
    };
  };

  govulncheck = final.buildGoModule rec {
    name = "govulncheck";
    version = "latest";
    src = final.fetchFromGitHub {
      owner = "golang";
      repo = "vuln";
      rev = "05fb7250142cc6010c39968839f2f3710afdd918";
      sha256 = "sha256-uxhcLnJftpQjBR26xAccoCmI+3AFdrTbySbSm2wchJs=";
    };
    vendorSha256 = "sha256-69v1jbqNTniDeLHAFJIq74e+zK979MdqsZ8OpcrDXbE=";

    doCheck = false;

    meta = with final.lib; {
      description = "the database client and tools for the Go vulnerability database";
      homepage = "https://github.com/golang/vuln";
      maintainers = [ "nhost" ];
      platforms = platforms.linux ++ platforms.darwin;
    };
  };


}
