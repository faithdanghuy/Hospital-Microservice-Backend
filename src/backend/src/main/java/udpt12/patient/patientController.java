package udpt12.patient;
import java.util.ArrayList;
import java.util.List;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class patientController {
    private List<Patient> patients = new ArrayList<>();

    @GetMapping("/patient")
    public List<Patient> findAll(){
        return patients;
    }
}
