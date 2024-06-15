{
  inputs = {
    fenix = {
      url = "github:nix-community/fenix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
    flake-utils.url = "github:numtide/flake-utils";
    nixpkgs.url = "nixpkgs/nixos-unstable";
  };

  outputs = { self, fenix, flake-utils, nixpkgs }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        toolchain = fenix.packages.${system}.minimal.toolchain;
        pkgs = nixpkgs.legacyPackages.${system};
      in
      {
        packages.default =
          (pkgs.makeRustPlatform {
            cargo = toolchain;
            rustc = toolchain;
          }).buildRustPackage {
            pname = "example";
            version = "0.1.0";

            src = ./.;

            cargoLock.lockFile = ./Cargo.lock;
          };

        devShells.default = mkShell {
          buildInputs = with pkgs; [
            pkg-config
            openssl
            wabt
            wasmtime
            wasm-tools
            toolchain
          ];
          PKG_CONFIG_PATH = "${pkgs.openssl.dev}/lib/pkgconfig";
          LD_LIBRARY_PATH = pkgs.lib.makeLibraryPath [ pkgs.openssl ];
        };
      }
    );
}
