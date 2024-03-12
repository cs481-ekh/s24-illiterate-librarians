import { Component } from '@angular/core';

@Component({
  selector: 'app-parent',
  templateUrl: './parent.component.html',
  styleUrl: './parent.component.css'
})
export class ProfileParentComponent {
  ngOnInit() {
    const editButton = document.getElementById('edit');
    const saveButton = document.getElementById('save');
    const profileEditing = document.getElementById('profile-editing');
    const profileView = document.getElementById('profile-view');
    

    if (editButton && saveButton && profileEditing && profileView) {
      editButton.addEventListener('click', function () {
        editButton.style.display = 'none';
        profileEditing.style.display = 'block';
        saveButton.style.display = 'block';
        profileView.style.display = 'none';
      });

      saveButton.addEventListener('click', function () {
        editButton.style.display = 'block';
        profileEditing.style.display = 'none';
        saveButton.style.display = 'none';
        profileView.style.display = 'block';
      });
    }
  }
}
