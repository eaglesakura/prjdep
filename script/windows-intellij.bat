@echo off

CD /d %~dp0\..\

SET GOPATH=%CD%\.gopath\windows
SET PATH=%MINGW64_PATH%\bin;%PATH%

%GO_IDEA_PATH%\bin\idea64.exe
