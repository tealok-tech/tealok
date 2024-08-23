{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.btrfs-progs
    pkgs.pkg-config
    pkgs.gnupg
    pkgs.go
    pkgs.gpgme
  ];
}
