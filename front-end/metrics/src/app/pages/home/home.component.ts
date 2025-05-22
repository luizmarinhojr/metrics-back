import { Component } from '@angular/core';
import { HeaderHomeComponent } from '../../components/header-home/header-home.component';
import { TableComponent } from "../../components/table/table.component";

@Component({
  selector: 'app-home',
  imports: [HeaderHomeComponent, TableComponent],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css'
})
export class HomeComponent {
  
}
