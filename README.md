# Stonks

Proyecto simple que simula un portafolio de acciones y realiza rebalanceos automáticos cuando cambian los precios.

## Visión general (alto nivel)

- Se modelan acciones (Stock) con precio y observadores.
- Se modela un portafolio (Portfolio) que contiene varias acciones, efectivo disponible y porcentajes objetivo por activo.
- El portafolio se suscribe a cambios de precio de cada acción. Cuando una acción cambia de precio notifica a sus observadores.
- Al recibir una notificación, el portafolio ejecuta un rebalanceo: compra o vende cantidades para aproximarse a las asignaciones objetivo, respetando el efectivo disponible.
- Hay utilidades para mostrar resumen y calcular el valor total del portafolio.

## Archivos clave

- domain/stock.go — definición de Stock y mecanismo de notificación a observadores.
- domain/portfolio.go — definición de Portfolio, lógica de rebalanceo y manejo de efectivo.
- domain/portfolio_test.go — pruebas unitarias que validan comportamiento y conservación de valor.
- main.go — ejemplo de uso / simulación.

## Ejecutar

- Ejecutar el ejemplo:
```sh
go run main.go
```

- Ejecutar tests:
```sh
go test ./...
```

- Generar reporte de cobertura (si existe Makefile):
```sh
make cover
```

## Comportamiento esperado

- El portafolio mantiene asignaciones por porcentaje (por ejemplo 60/40).
- Cuando sube o baja el precio de una acción, el portafolio intenta ajustar posiciones para volver a las proporciones objetivo, usando el efectivo disponible.
- Las pruebas verifican que el valor total del portafolio y las proporciones se mantienen razonablemente tras cambios de precio.

## Notas

- Proyecto pensado como ejemplo educativo; no es software financiero para producción.