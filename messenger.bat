@echo off
color Ob
title Chat Messenger
echo Hi, welcome to the chat program!
echo.
set /p username="Enter Your Name: "
echo %username% has joined! >> join.dat

:message
set /p message="Say: "
goto send

:send
echo %username%: %message% (%time%)>> log.dat
set /p do="Leave Chat?"
if "%do%"=="yes" (exit) else (goto message)