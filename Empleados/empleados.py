import psycopg2
from psycopg2 import sql

config = {
    'host': 'localhost',
    'database': 'empleados',
    'user': 'postgres',
    'password': '3530', 
    'port': '5432'
}

def conectar():
    """Establece conexión con la base de datos"""
    try:
        conn = psycopg2.connect(**config)
        return conn
    except Exception as e:
        print(f"Error al conectar a PostgreSQL: {e}")
        return None

def crear_tabla(conn):
    """Crea la tabla empleados si no existe"""
    try:
        with conn.cursor() as cur:
            cur.execute("""
                CREATE TABLE IF NOT EXISTS empleados (
                    id SERIAL PRIMARY KEY,
                    nombre VARCHAR(100) NOT NULL,
                    cargo VARCHAR(100) NOT NULL,
                    salario DECIMAL(10, 2) NOT NULL
                )
            """)
            conn.commit()
            print("Tabla 'empleados' creada exitosamente")
    except Exception as e:
        print(f"Error al crear tabla: {e}")
        conn.rollback()

def insertar_datos(conn):
    """Inserta datos de ejemplo en la tabla"""
    empleados = [
        ('Juan Pérez', 'Gerente', 3500000.00),
        ('María Gómez', 'Desarrollador Senior', 2800000.00),
        ('Carlos Rojas', 'Analista de Datos', 1900000.00),
        ('Ana López', 'Desarrollador Junior', 2200000.00),
        ('Pedro Sánchez', 'Diseñador', 2400000.00)
    ]
    
    try:
        with conn.cursor() as cur:
            cur.executemany(
                "INSERT INTO empleados (nombre, cargo, salario) VALUES (%s, %s, %s)",
                empleados
            )
            conn.commit()
            print(f"{len(empleados)} registros insertados correctamente")
    except Exception as e:
        print(f"Error al insertar datos: {e}")
        conn.rollback()

def consultar_salarios_altos(conn):
    """Consulta empleados con salario > 2 millones"""
    try:
        with conn.cursor() as cur:
            cur.execute("""
                SELECT id, nombre, cargo, salario
                FROM empleados
                WHERE salario > 2000000
                ORDER BY salario DESC
            """)
            
            print("\nEmpleados con salario mayor a 2 millones:")
            print("-"*50)
            print("{:<5} {:<20} {:<25} {:<10}".format(
                "ID", "Nombre", "Cargo", "Salario"))
            print("-"*50)
            
            for empleado in cur.fetchall():
                print("{:<5} {:<20} {:<25} ${:<10,.2f}".format(
                    empleado[0], empleado[1], empleado[2], empleado[3]))
                
    except Exception as e:
        print(f"Error al consultar datos: {e}")

def main():
    # contexión
    conn = conectar()
    if not conn:
        return
    
    try:
        # Crear tabla
        crear_tabla(conn)
        
        # Insertar datos
        insertar_datos(conn)
        
        # Consultar salarios altos
        consultar_salarios_altos(conn)
        
    finally:
        # Cerrar conexión
        if conn:
            conn.close()
            print("\nConexión cerrada")

if __name__ == "__main__":
    main()