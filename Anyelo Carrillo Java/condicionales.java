import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        System.out.print("Seleccione la figura a calcular su área:\n\n 1 para rombo\n\n 2 para rectángulo\n\n 3 para cuadrado\n\n 4 para trapecio\n\n: ");
        String Figura = scanner.nextLine();

        double Pi = 3.1416;
        int const = 2;

        switch (Figura) {
            case "1":
                System.out.print("Ingresa el valor de la diagonal mayor: ");
                int Dmayor = scanner.nextInt();
                System.out.print("Ingresa el valor de la diagonal menor: ");
                int Dmenor = scanner.nextInt();
                double AreaRombo = (Dmayor + Dmenor) / const;
                System.out.println("El área del rombo es: " + AreaRombo);
                break;
            case "2":
                System.out.print("Ingresa el valor del largo del rectángulo: ");
                int Largo = scanner.nextInt();
                System.out.print("Ingresa el valor del ancho del rectángulo: ");
                int Ancho = scanner.nextInt();
                int AreaRectangulo = Largo * Ancho;
                System.out.println("El área del rectángulo es: " + AreaRectangulo);
                break;
            case "3":
                System.out.print("Ingresa el valor del lado del cuadrado: ");
                int Lado = scanner.nextInt();
                int AreaCuadrado = Lado * Lado;
                System.out.println("El área del cuadrado es: " + AreaCuadrado);
                break;
            case "4":
                System.out.print("Ingresa el valor de la base mayor: ");
                int Bmayor = scanner.nextInt();
                System.out.print("Ingresa el valor de la base menor: ");
                int Bmenor = scanner.nextInt();
                System.out.print("Ingrese la altura del trapecio: ");
                int H = scanner.nextInt();
                double AreaTrapecio = ((Bmayor + Bmenor) * H) / 2.0;
                System.out.println("El área del trapecio es: " + AreaTrapecio);
                break;
            default:
                System.out.println("Opción no válida.");
        }
    }
}
