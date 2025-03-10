import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        // Crear un objeto Scanner para leer la entrada del usuario
        Scanner scanner = new Scanner(System.in);

        // Solicitar al usuario que ingrese el primer valor
        System.out.print("Ingrese el primer valor: ");
        int A = scanner.nextInt(); // Leer el primer valor como entero

        // Inicializar la variable B (aunque no se usa en el código)
        int B = 0;

        // Solicitar al usuario que ingrese el segundo valor
        System.out.print("Ingrese el segundo valor: ");
        int C = scanner.nextInt(); // Leer el segundo valor como entero

        // Calcular la potencia de A elevado a C
        double valor = Math.pow(A, C); // Usar Math.pow para calcular la potencia

        // Mostrar el resultado
        System.out.println("La potencia de " + A + " sobre " + C + " es: " + valor);

        // Cerrar el Scanner
        scanner.close();
    }
}
