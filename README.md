# Sistema de Gestión de Libros Electrónicos

## Descripción General

El **Sistema de Gestión de Libros Electrónicos** es una aplicación desarrollada en **Golang** que permite administrar un catálogo digital de libros electrónicos, facilitando la gestión de autores, categorías, usuarios y registros de descargas.

El proyecto ha sido concebido como una implementación práctica de los principios de **Programación Orientada a Objetos (POO)** aplicados en Go, combinando técnicas de modelado UML, persistencia de datos mediante **MySQL** y una arquitectura organizada por capas.

El sistema permite a los usuarios consultar información bibliográfica, explorar el catálogo disponible y registrar descargas de libros electrónicos, mientras que los administradores pueden gestionar el contenido del catálogo mediante operaciones de creación, edición y eliminación de registros.

---

## Objetivo del Proyecto

Diseñar e implementar una solución orientada a objetos para la administración de libros electrónicos, aplicando conceptos fundamentales de ingeniería de software, modelado UML, encapsulación de datos, abstracción, composición, persistencia y manejo de errores.

---

## Características Principales

### Gestión de Libros

- Registro de libros electrónicos.
- Consulta de información bibliográfica.
- Actualización de información de libros.
- Eliminación de libros del catálogo.
- Control de disponibilidad.

### Gestión de Autores

- Registro de autores.
- Asociación de autores a libros.
- Consulta de información relacionada.

### Gestión de Categorías

- Clasificación de libros por categorías.
- Organización del catálogo digital.

### Gestión de Usuarios

- Administración de lectores.
- Administración de usuarios con privilegios administrativos.

### Gestión de Descargas

- Registro histórico de descargas realizadas.
- Seguimiento de usuario, libro y fecha de descarga.

### Persistencia de Datos

- Almacenamiento de información mediante MySQL.
- Acceso a datos mediante el patrón Repository.

---

## Principios de Programación Orientada a Objetos Aplicados

### Encapsulamiento

Las entidades del dominio utilizan atributos privados y exponen únicamente los métodos necesarios para acceder o modificar la información.

**Ejemplo de entidades:**

- Libro
- Usuario
- Autor
- Categoría
- Descarga

---

### Abstracción

Se utilizan interfaces para definir comportamientos comunes sin depender de implementaciones concretas.

**Ejemplos:**

- `Descargable`
- `LibroRepository`

---

### Polimorfismo

Las interfaces permiten que distintas implementaciones compartan un mismo contrato funcional, favoreciendo la extensibilidad y el desacoplamiento de componentes.

---

### Composición

Go favorece la composición sobre la herencia tradicional.

**Modelo conceptual:**

```text
Usuario
├── Lector
└── Administrador
```

Implementado mediante *embedding* de estructuras.

---

### Separación de Responsabilidades

Cada capa del sistema posee una función claramente definida:

- Domain
- Services
- Repository
- Configuration

---

## Modelo del Dominio

El sistema se encuentra compuesto por las siguientes entidades principales:

```text
Usuario
├── Lector
└── Administrador

Biblioteca
Libro
Autor
Categoria
Descarga
```

Estas entidades fueron identificadas mediante análisis orientado a objetos y representadas mediante diagramas UML.

---

## Arquitectura del Proyecto

```text
ebook-management-system/
│
├── cmd/
│   └── app/
│       └── main.go
│
├── internal/
│   ├── config/
│   ├── domain/
│   ├── interfaces/
│   ├── repository/
│   ├── services/
│   └── errors/
│
├── docs/
│   └── uml/
│
├── scripts/
│
└── README.md
```

### Domain

Contiene las entidades y reglas centrales del negocio.

### Interfaces

Define contratos y comportamientos abstractos.

### Repository

Implementa el acceso a datos utilizando MySQL.

### Services

Contiene la lógica de negocio de la aplicación.

### Config

Gestiona la configuración y conexión con la base de datos.

---

## Tecnologías Utilizadas

- Go (Golang)
- MySQL
- Mermaid
- UML
- Git
- GitHub

---

## Base de Datos

La aplicación utiliza MySQL como motor de persistencia para almacenar:

- Libros
- Autores
- Categorías
- Usuarios
- Descargas

El acceso a los datos se realiza mediante el driver:

```bash
github.com/go-sql-driver/mysql
```

---

## Diagramas UML

El proyecto incluye documentación UML para facilitar el análisis, diseño e implementación del sistema.

### Diagramas Incluidos

- Diagrama de Clases
- Diagrama de Casos de Uso
- Diagramas de Secuencia
- Modelo Relacional de Base de Datos

---

## Buenas Prácticas Implementadas

- Arquitectura por capas.
- Separación de responsabilidades.
- Principios SOLID.
- Manejo centralizado de errores.
- Uso de interfaces para desacoplamiento.
- Encapsulamiento mediante atributos privados.
- Persistencia desacoplada mediante repositorios.
- Control de versiones con Git y GitHub.

---

## Objetivos Académicos Alcanzados

- Aplicación práctica de Programación Orientada a Objetos en Golang.
- Diseño de software basado en UML.
- Implementación de Encapsulamiento, Abstracción, Composición y Polimorfismo.
- Integración de MySQL como mecanismo de persistencia.
- Implementación de una arquitectura escalable y mantenible.
- Aplicación de principios de diseño orientado a objetos y SOLID.
- Documentación técnica y modelado previo al desarrollo.

---

## Ejecución del Proyecto

### Clonar el repositorio

```bash
git clone https://github.com/usuario/ebook-management-system.git
```

### Ingresar al proyecto

```bash
cd ebook-management-system
```

### Instalar dependencias

```bash
go mod tidy
```

### Configurar la base de datos

```bash
mysql -u root -p < scripts/schema.sql
```

### Ejecutar la aplicación

```bash
go run cmd/app/main.go
```

---

## Autor

**Nelson Cacoango**

Proyecto académico desarrollado como aplicación práctica de Programación Orientada a Objetos utilizando **Golang** y **MySQL** para la gestión de libros electrónicos.
