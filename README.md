# ProteiTestCase
For start:
1. systemd --user enable go-server-protei && \
   systemd --user start go-server-protei && \
   systemd --user status go-server-protei
2. sudo nano /etc/systemd/logind.conf
3. Write in last line: UserStopDelaySec=infinity
4. make server-build
   make deploy
   make restart-service
