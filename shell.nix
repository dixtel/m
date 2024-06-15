let pkgs = import <nixpkgs> { }; in
pkgs.mkShell {
  buildInputs = with pkgs; [
    pkg-config
    openssl
  ];
  PKG_CONFIG_PATH = "${pkgs.openssl.dev}/lib/pkgconfig";
  LD_LIBRARY_PATH = pkgs.lib.makeLibraryPath [ pkgs.openssl ];
}
