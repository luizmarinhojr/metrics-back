import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-add',
  imports: [],
  templateUrl: './add.component.html',
  styleUrl: './add.component.css'
})
export class AddComponent implements OnInit{
  date!: string;

  constructor() {}

  ngOnInit(): void {
    this.setCurrentDate();
  }

  setCurrentDate(): void {
    const hoje = new Date();
    this.date = hoje.toISOString().substring(0, 10);
  }
}
