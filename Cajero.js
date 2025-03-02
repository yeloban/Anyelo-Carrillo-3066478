const readline = require("readline"); // Importamos el módulo readline para manejar la entrada del usuario en la consola

// Clase Usuario que representa a un cliente del banco
class Usuario {
    constructor(nombre, pin, saldo) {
        this.nombre = nombre; // Nombre del usuario
        this.pin = pin; // PIN de acceso
        this.saldo = saldo; // Saldo en la cuenta
    }
}

// Clase Banco que maneja los usuarios y sus operaciones
class Banco {
    constructor(usuarios = []) {
        this.usuarios = usuarios; // Lista de usuarios registrados en el banco
    }

    // Método para autenticar al usuario
    autenticar(nombre, pin) {
        for (let usuario of this.usuarios) {
            if (usuario.nombre === nombre) { // Comprobamos si el usuario existe
                if (usuario.pin === pin) { // Verificamos si el PIN es correcto
                    console.log("¡Estás logueado!");
                    return true;
                } else {
                    console.log("Pin o nombre incorrecto.");
                    return false;
                }
            }
        }
        console.log("El usuario no existe.");
        return false;
    }

    // Método para retirar dinero de la cuenta de un usuario
    sacarDinero(nombre, cantidad) {
        for (let usuario of this.usuarios) {
            if (usuario.nombre === nombre) {
                if (usuario.saldo < cantidad) { // Verificamos si el saldo es suficiente
                    console.log("Saldo insuficiente.");
                } else {
                    usuario.saldo -= cantidad; // Restamos la cantidad al saldo
                    console.log("El saldo disponible es: " + usuario.saldo);
                }
                return;
            }
        }
    }
}

// Crear instancias de usuarios con nombre, PIN y saldo inicial
const ana = new Usuario("Ana", 9872, 450);
const pablo = new Usuario("Pablo", 5555, 200);
const rodrigo = new Usuario("Rodrigo", 2222, 900);

// Crear una instancia del Banco con los usuarios registrados
const banco = new Banco([ana, pablo, rodrigo]);

// Configuración del readline para manejar la entrada y salida de datos en consola
const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

// Función para iniciar sesión en el sistema
function iniciarSesion() {
    rl.question("Introduzca su nombre: ", (nombre) => {
        rl.question("Introduzca su pin: ", (pin) => {
            pin = parseInt(pin); // Convertimos el PIN a número
            if (banco.autenticar(nombre, pin)) { // Si la autenticación es exitosa, mostramos el menú
                menuUsuario(nombre);
            } else {
                iniciarSesion(); // Si falla, pedimos los datos nuevamente
            }
        });
    });
}

// Función que muestra el menú de opciones para el usuario autenticado
function menuUsuario(nombre) {
    rl.question("Por favor, elija una opción:\n1. Sacar dinero\n2. Terminar sesión\n", (opcion) => {
        if (opcion === "1") { // Opción para retirar dinero
            rl.question("Introduce cantidad a sacar: ", (cantidad) => {
                banco.sacarDinero(nombre, parseInt(cantidad)); // Llamamos al método para retirar dinero
                menuUsuario(nombre); // Volvemos a mostrar el menú
            });
        } else if (opcion === "2") { // Opción para cerrar sesión
            console.log("Saliendo del sistema.");
            rl.close(); // Cerramos la interfaz readline
        } else { // Opción inválida
            console.log("Opción incorrecta.");
            menuUsuario(nombre); // Volvemos a mostrar el menú
        }
    });
}

// Mensaje de bienvenida e inicio de sesión
console.log("Bienvenido al Banco");
iniciarSesion();
