public class Main {
    public static void main(String[] args) {
        int a = 2;
        if (a == 2) {
            System.out.println("a vale 2");
        } else {
            System.out.println("a es diferente de 2");
        }

        String Nombre = "Esteban";
        int Edad = 28;
        String Pais = "Colombia";

        if (Nombre.equals("Esteban") && Edad == 28 && Pais.equals("Colombia")) {
            System.out.println("Su nombre es " + Nombre + ", tiene " + Edad + " años y es de " + Pais);
        } else if (Nombre.equals("Esteban") && Edad != 28) {
            System.out.println("Su nombre es Esteban y no tiene 28 años");
        } else if (!Nombre.equals("Esteban") && Edad == 28) {
            System.out.println("Su nombre no es Esteban y tiene 28 años");
        } else {
            System.out.println("No se llama Esteban ni tiene 28 años");
        }
    }
}

         