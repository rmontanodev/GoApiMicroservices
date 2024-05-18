#!/bin/bash

# Ejecuta las pruebas y genera el perfil de cobertura
go test ./... -coverprofile=coverage.out

# Genera el informe de cobertura en HTML
go tool cover -html=coverage.out -o coverage.html

# Muestra el informe en el navegador (opcional)
# open coverage.html # macOS
# xdg-open coverage.html # Linux
