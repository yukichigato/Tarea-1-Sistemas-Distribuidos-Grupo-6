package utils

const (
	MENU_HOME_TEXT = `
	Menu

	1. Registrarse
	2. Iniciar sesión
	3. Terminar ejecución
	`

	MENU_PRINCIPAL_TEXT = `
	Menu
	
	1. Ver catálogo
	2. Carros de compras
	3. Mis prestamos
	4. Mi cuenta
	5. Populares
	6. Salir

	Seleccione una opción:
	`

	MENU_MI_CUENTA_TEXT = `
	1. Consultar saldo
	2. Abonar usm pesos
	3. Ver historial de compras y arriendos
	4. Salir
	`

	CATALOGO_TABLA_HEADER = `	
	------------------------------------------------------------------------
	| ID | Título               | Categoría          | Modalidad   | Valor |
	------------------------------------------------------------------------
	`

	CATALOGO_TABLE_FOOTER = `
	------------------------------------------------------------------------
	`

	CARRO_COMPRA_TABLE_HEADER = `
	---------------------------------------------------------------------------------
	| Título               | Modalidad          | Valor       | Fecha de devolución |
	---------------------------------------------------------------------------------
	`

	CARRO_COMPRA_TABLE_FOOTER = `
	---------------------------------------------------------------------------------
	`

	PRESTAMOS_TABLE_HEADER = `
	----------------------------------------------------------------------------------------------------
	| ID Préstamo | Título               | Fecha de inicio | Fecha de fin | Días restantes | Estado    |
	----------------------------------------------------------------------------------------------------
	`

	PRESTAMOS_TABLE_FOOTER = `
	----------------------------------------------------------------------------------------------------
	`

	HISTORIAL_COMPRAS_ARRIENDOS_TABLE_HEADER = `
	----------------------------------------------------------------------------------------------------- 
	| ID Transacción | ID Libro | Título               | Tipo        | Fecha de transacción | Valor     |
	----------------------------------------------------------------------------------------------------- 
	`

	HISTORIAL_COMPRAS_ARRIENDOS_TABLE_FOOTER = `
	----------------------------------------------------------------------------------------------------- 
	`

	POPULARES_TABLE_HEADER = `
	----------------------------------------------------------------
	| ID Libro | Título               | Categoría    | Popularidad |
	----------------------------------------------------------------
	`

	POPULARES_TABLE_FOOTER = `
	----------------------------------------------------------------
	`
)