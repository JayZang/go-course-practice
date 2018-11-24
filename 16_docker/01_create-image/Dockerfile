FROM docker/whalesay:latest
RUN apt-get -y update && apt-get install -y fortunes
CMD pwd; ls -la && /usr/games/fortune -a | cowsay
