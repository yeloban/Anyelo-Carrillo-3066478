import java.util.InputMismatchException;
import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        // Parte 1: Bucle while con contador
        int contador = 0;
        while (contador < 30) {
            contador++;
            System.out.println("Iteración " + contador);
        }

        // Parte 2: Bucle while con if-else y manejo de errores
        Scanner scanner = new Scanner(System.in);
        while (true) {
            try {
                System.out.print("Introduzca un valor mayor a 10: ");
                int a = scanner.nextInt(); // Lee un entero desde la entrada

                if (a > 10) {
                    System.out.println("Es correcto");
                    break; // Sale del bucle cuando la condición se cumple
                } else {
                    System.out.println("Es incorrecto, pruebe de nuevo");
                }
            } catch (InputMismatchException e) {
                System.out.println("Por favor, ingrese un número válido");
                scanner.next(); // Limpia el buffer de entrada
            }
        }
        scanner.close(); // Cierra el scanner
    }
}