import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        // Crear un objeto Scanner para leer la entrada del usuario
        Scanner scanner = new Scanner(System.in);

        // Solicitar y leer el valor del voltaje
        System.out.print("Ingrese el valor del voltaje del circuito: ");
        double Voltaje = scanner.nextDouble();

        // Solicitar y leer el valor de la resistencia
        System.out.print("Ingrese el valor de la resistencia del circuito: ");
        double Resistencia = scanner.nextDouble();

        // Calcular la intensidad (corriente)
        double Intensidad = Voltaje / Resistencia;

        // Mostrar el resultado
        System.out.println("Al conectar un resistor de R " + Resistencia + " ohm a una fuente de V " + Voltaje + " voltios, circulará una corriente de " + Intensidad + " amperios");

        // Cerrar el Scanner
        scanner.close();
    }
}
