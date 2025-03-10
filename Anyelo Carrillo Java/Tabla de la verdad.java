public class Tabla_de_verdad {
    public static void main(String[] args) {
        // Parte 1: Operaciones con booleanos
        boolean a = true;
        boolean b = false;
        System.out.println(a && b); // Imprime "false"

        // Parte 2: Operaciones con enteros y comparaciones
        int aNum = 2;
        int bNum = 3;
        int c = 4;
        int d = 10;

        // Comparación 1: a == b y c < d
        System.out.println((aNum == bNum) && (c < d)); // Imprime "false"

        // Comparación 2: not (a == b) y c > d
        System.out.println(!(aNum == bNum) && (c > d)); // Imprime "false"
    }
}
