{
  description = "Command executor and theming tool";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs = { self, nixpkgs }:
  let
    supportedSystems = [ "x86_64-linux" "aarch64-linux" "x86_64-darwin" "aarch64-darwin" ];

    # Helper function to generate attributes for all systems
    forEachSupportedSystem = f: nixpkgs.lib.genAttrs supportedSystems (system: f {
      pkgs = import nixpkgs { inherit system; };
    });
  in
  {
    packages = forEachSupportedSystem( { pkgs }: {
      default = pkgs.buildGoModule {
        pname = "barista";
        version = "0.1";

        src = "./.";
        vendorHash = pkgs.lib.fakeHash;

        meta = with pkgs.lib; {
          description = "A theming utillity and command executor written in go";
          homepage = "https://github.com/Lucky44x/barista";
          license = licenses.mit;
          maintainers = [ Lucky44x ];
        };
      };
    });

    devShells = forEachSupportedSystem ({ pkgs }: {
      default = pkgs.mkShell {
        packages = with pkgs; [ go gopls ];
      };
    });
  };
}
