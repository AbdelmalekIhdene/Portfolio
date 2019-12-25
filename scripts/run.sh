mkdir -p ~/run
nohup ./portfolio &> /dev/null & echo $! > ~/run/portfolio.pid
