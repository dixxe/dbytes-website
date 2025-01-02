{
  description = "A portable go dev flake!";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
  };

  outputs = { nixpkgs, ... }@inputs: 
	let
    system = "x86_64-linux";
		pkgs = nixpkgs.legacyPackages.${system};
  in 
	{
    devShells.${system} = {
		  default = pkgs.mkShell {
			  nativeBuildInputs = with pkgs; [
				  go
					gopls
					templ
					bashInteractive
				];
				
				hardeningDisable = [ "fortify" ];

				shellHook = "echo Welcome to go-backend environment";
			};
		};
	};
}
