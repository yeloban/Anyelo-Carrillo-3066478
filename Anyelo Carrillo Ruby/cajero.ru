# Clase Usuario que representa a un cliente del banco
class Usuario
    attr_accessor :nombre, :pin, :saldo  # Permite acceder y modificar los atributos
  
    def initialize(nombre, pin, saldo)
      @nombre = nombre  # Nombre del usuario
      @pin = pin        # PIN de seguridad
      @saldo = saldo    # Saldo disponible
    end
  end
  
  # Clase Banco que maneja los usuarios y transacciones
  class Banco
    def initialize(usuarios = [])  # Se inicializa con una lista de usuarios
      @usuarios = usuarios
    end
  
    # Método para autenticar al usuario por nombre y PIN
    def autenticar(nombre, pin)
      @usuarios.each do |usuario|
        if usuario.nombre == nombre
          if usuario.pin == pin
            puts "¡Estas logeado!"
            return true
          else
            puts "Pin o nombre incorrecto"
            return false
          end
        end
      end
      puts "El usuario no existe"
      false
    end
  
    # Método para retirar dinero del saldo de un usuario
    def sacar_dinero(nombre, cantidad)
      @usuarios.each do |usuario|
        if usuario.nombre == nombre
          if usuario.saldo < cantidad
            puts "Saldo insuficiente"
          else
            usuario.saldo -= cantidad
            puts "El saldo disponible es: #{usuario.saldo}"
          end
        end
      end
    end
  end
  
  # Creación de usuarios
  ana = Usuario.new('Ana', 9872, 450)
  pablo = Usuario.new('Pablo', 5555, 200)
  rodrigo = Usuario.new('Rodrigo', 2222, 900)
  
  # Creación del banco con los usuarios
  banco = Banco.new([ana, pablo, rodrigo])
  
  # Bucle principal del sistema de cajero automático
  loop do
    puts "Bienvenido al Banco, por favor, identifíquese."
    print "Introduzca el nombre: "
    nombre = gets.chomp  # chomp elimina el salto de línea
  
    print "Introduzca el PIN: "
    pin = gets.chomp.to_i  # Convertimos la entrada a número entero
  
    if banco.autenticar(nombre, pin)
      loop do
        puts "Por favor, elija una de estas opciones:"
        puts "1. Sacar dinero"
        puts "2. Terminar sesión"
        print "Opción: "
        opcion = gets.chomp
  
        case opcion
        when '1'
          print "Introduce cantidad a sacar: "
          cantidad = gets.chomp.to_i
          banco.sacar_dinero(nombre, cantidad)
        when '2'
          puts "Saliendo del sistema..."
          break
        else
          puts "Opción incorrecta. Por favor, introduzca otra opción."
        end
      end
    else
      puts "Usuario no autenticado. Introduzca nombre y PIN correctos."
    end
  end
  