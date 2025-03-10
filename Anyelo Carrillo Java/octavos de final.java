import java.util.ArrayList;
import java.util.List;
import java.util.Random;
import java.util.Scanner;

public class torneo {

    // Función para jugar un partido
    static String jugarPartido(String equipo1, String equipo2, Scanner scanner) {
        System.out.println("\n" + equipo1 + " vs " + equipo2);

        System.out.print("Ingrese goles de " + equipo1 + ": ");
        int goles1 = scanner.nextInt();

        System.out.print("Ingrese goles de " + equipo2 + ": ");
        int goles2 = scanner.nextInt();

        if (goles1 > goles2) {
            System.out.println(" " + equipo1 + " gana y avanza a la siguiente ronda!");
            return equipo1;
        } else if (goles2 > goles1) {
            System.out.println(" " + equipo2 + " gana y avanza a la siguiente ronda!");
            return equipo2;
        } else {
            System.out.println(" EMPATE! Se jugará tiempo extra...");
            return jugarTiempoExtra(equipo1, equipo2, scanner);
        }
    }

    // Función para jugar tiempo extra
    static String jugarTiempoExtra(String equipo1, String equipo2, Scanner scanner) {
        System.out.print("Tiempo extra - Goles de " + equipo1 + ": ");
        int goles1 = scanner.nextInt();

        System.out.print("Tiempo extra - Goles de " + equipo2 + ": ");
        int goles2 = scanner.nextInt();

        if (goles1 > goles2) {
            System.out.println(" " + equipo1 + " gana en tiempo extra!");
            return equipo1;
        } else if (goles2 > goles1) {
            System.out.println(" " + equipo2 + " gana en tiempo extra!");
            return equipo2;
        } else {
            System.out.println(" ¡Sigue el empate! Se define por penales...");
            return jugarPenales(equipo1, equipo2);
        }
    }

    // Función para jugar penales
    static String jugarPenales(String equipo1, String equipo2) {
        Random random = new Random();
        while (true) {
            int penales1 = random.nextInt(5) + 1;
            int penales2 = random.nextInt(5) + 1;

            System.out.println(" Penales: " + equipo1 + " " + penales1 + " - " + penales2 + " " + equipo2);

            if (penales1 > penales2) {
                System.out.println(" " + equipo1 + " gana por penales!");
                return equipo1;
            } else if (penales2 > penales1) {
                System.out.println(" " + equipo2 + " gana por penales!");
                return equipo2;
            } else {
                System.out.println(" Sigue el empate en penales... ¡Se repiten los tiros!");
            }
        }
    }

    // Función para simular una ronda
    static List<String> simularRonda(List<String> equipos, String nombreRonda, Scanner scanner) {
        System.out.println("\n===== " + nombreRonda + " =====");
        List<String> ganadores = new ArrayList<>();

        for (int i = 0; i < equipos.size(); i += 2) {
            String equipo1 = equipos.get(i);
            String equipo2 = equipos.get(i + 1);
            String ganador = jugarPartido(equipo1, equipo2, scanner);
            ganadores.add(ganador);
        }

        return ganadores;
    }

    public static void main(String[] args) {
        // Lista de equipos en octavos de final
        List<String> equipos = new ArrayList<>(List.of(
            "Brasil", "Corea del Sur", "Francia", "Dinamarca",
            "Argentina", "México", "España", "Alemania",
            "Japón", "Inglaterra", "Perú", "USA",
            "Nigeria", "Chile", "Portugal", "Colombia"
        ));

        // Nombres de las rondas
        String[] rondas = {"Octavos de Final", "Cuartos de Final", "Semifinal", "Final"};

        // Crear un objeto Scanner para leer la entrada del usuario
        Scanner scanner = new Scanner(System.in);

        // Simulación del torneo
        List<String> rondaActual = equipos;
        for (String nombreRonda : rondas) {
            rondaActual = simularRonda(rondaActual, nombreRonda, scanner);
        }

        // Anunciar al campeón
        System.out.println("\n  EL CAMPEÓN DEL MUNDIAL ES: " + rondaActual.get(0) + " ");

        // Cerrar el Scanner
        scanner.close();
    }
}
