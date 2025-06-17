package udpt12.patient;

public class Patient {
    private Long id;
    private String full_name;
    private Integer age;
    private String gender;
    private String phone_number;
    private String email;

    public Patient(Long id, String full_name, Integer age, String gender, String phone_number, String email) {
        this.id = id;
        this.full_name = full_name;
        this.age = age;
        this.gender = gender;
        this.phone_number = phone_number;
        this.email = email;
    }
}
