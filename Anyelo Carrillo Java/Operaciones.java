import java.util.Scanner;

public class Operaciones {
    public static void main(String[] args) {
        // Crear un objeto Scanner para leer la entrada del usuario
        Scanner scanner = new Scanner(System.in);

        // Solicitar y leer el primer valor
        System.out.print("Ingrese valor: ");
        int A = scanner.nextInt();

        // Solicitar y leer el segundo valor
        System.out.print("Ingrese valor: ");
        int B = scanner.nextInt();

        // Realizar operaciones
        int suma = A + B;
        int res = A - B;
        int multi = A * B;
        double div = (double) A / B; // Conversión a double para división precisa

        // Mostrar resultados de las operaciones
        System.out.println("La suma de los números es: " + suma);
        System.out.println("La resta de los números es: " + res);
        System.out.println("La multiplicación de los números es: " + multi);
        System.out.println("La división de los números es: " + div);

        // Comparaciones
        boolean igual = (A == B);
        boolean mayor = (A > B);

        // Mostrar resultados de las comparaciones
        System.out.println("El número es igual: " + igual);
        System.out.println("El número menor es: " + (mayor ? B : A)); // Operador ternario
        System.out.println("El número mayor es: " + (mayor ? A : B)); // Operador ternario

        // Cerrar el Scanner
        scanner.close();
    }
}
