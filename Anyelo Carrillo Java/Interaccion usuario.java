import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        // Crear un objeto Scanner para leer la entrada del usuario
        Scanner scanner = new Scanner(System.in);

        // Solicitar y leer los nombres completos
        System.out.print("Escriba sus nombres completos: ");
        String a = scanner.nextLine();

        // Solicitar y leer los apellidos completos
        System.out.print("Escriba sus apellidos completos: ");
        String b = scanner.nextLine();

        // Solicitar y leer la profesión
        System.out.print("Escriba su profesión: ");
        String c = scanner.nextLine();

        // Solicitar y leer el año de nacimiento
        System.out.print("Escriba su año de nacimiento: ");
        int d = scanner.nextInt();

        // Calcular la edad en el año 2025
        int e = 2025 - d;

        // Mostrar el resultado
        System.out.println("El (La) " + c + " " + a + " " + b + " tiene " + e + " años");

        // Cerrar el Scanner
        scanner.close();
    }
}
