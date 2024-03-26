import { Component } from '@angular/core';

@Component({
  selector: 'app-children',
  templateUrl: './children.component.html',
  styleUrl: './children.component.css'
})
export class ProfileChildrenComponent {
  ngOnInit() {
    const childBoxes = document.querySelectorAll('.child-box');
    const childModal = document.getElementById('childModal');
    const childClose = document.getElementById('childClose');
    const addButton = document.getElementById('addButton');
    const addModal = document.getElementById('addModal');
    const addClose = document.getElementById('addClose');


    if (childBoxes && childClose && childModal && addButton && addModal && addClose) {
      childBoxes.forEach(childBox => {
        childBox.addEventListener('click', function () {
          childModal.style.display = 'block';
        });
      });

      childClose.addEventListener('click', function () {
        childModal.style.display = 'none';
      });

      addButton.addEventListener('click', function () {
        addModal.style.display = 'block';
      });

      addClose.addEventListener('click', function () {
        addModal.style.display = 'none';
      });
    }
  }
}
