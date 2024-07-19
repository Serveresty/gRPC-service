# gRPC service (ProteiTestCase)
<p>Сервис с авторизацией, который реализует функции получения информации о сотруднике, а так же даты его отсутствия и причины.</p>
For start:
<ol>
<li> systemd --user enable go-server-protei && \
     systemd --user start go-server-protei && \
     systemd --user status go-server-protei
</li>
<li> sudo nano /etc/systemd/logind.conf </li>
<li> Write in last line: UserStopDelaySec=infinity </li>
<li> make server-build
     make deploy
     make restart-service
</li>
</ol>
