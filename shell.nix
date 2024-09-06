{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.btrfs-progs
    pkgs.gnupg
    pkgs.go
    pkgs.go-bindata
    pkgs.go-migrate
    pkgs.gpgme
    pkgs.pkg-config
    pkgs.sqlc
  ];
}
