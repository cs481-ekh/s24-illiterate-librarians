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


    if (childBoxes && childClose && childModal) {
      childBoxes.forEach(childBox => {
        childBox.addEventListener('click', function () {
          childModal.style.display = 'block';
        });
      });

      childClose.addEventListener('click', function () {
        childModal.style.display = 'none';
      });
    }
  }
}
