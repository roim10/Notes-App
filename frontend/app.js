const notesContainer = document.getElementById("notes");
const titleInput = document.getElementById("title");
const contentInput = document.getElementById("content");

// Загрузка всех заметок
async function fetchNotes() {
  const res = await fetch("http://localhost:8080/notes");
  const notes = await res.json();
  notesContainer.innerHTML = "";
  notes.forEach((note) => {
    const noteElem = document.createElement("div");
    noteElem.className = "note show";
    noteElem.id = `note-${note.id}`;
    noteElem.innerHTML = `
      <strong>${note.title}</strong>
      <p>${note.content}</p>
      <div class="buttons">
        <button onclick="editNote(${note.id})">✏️</button>
        <button onclick="deleteNote(${note.id})">🗑️</button>
      </div>
    `;
    notesContainer.appendChild(noteElem);
  });
}

// Добавление заметки
async function addNote() {
  if (!titleInput.value || !contentInput.value) return;
  await fetch("http://localhost:8080/notes", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      title: titleInput.value,
      content: contentInput.value,
    }),
  });
  titleInput.value = "";
  contentInput.value = "";
  fetchNotes();
}

// Удаление с анимацией растворения
async function deleteNote(id) {
  const noteElem = document.getElementById(`note-${id}`);
  noteElem.classList.add("hide");
  setTimeout(async () => {
    await fetch(`http://localhost:8080/notes/${id}`, { method: "DELETE" });
    noteElem.remove();
  }, 500);
}

// Редактирование заметки
function editNote(id) {
  const note = document.getElementById(`note-${id}`);
  const newTitle = prompt(
    "Новый заголовок:",
    note.querySelector("strong").textContent,
  );
  const newContent = prompt(
    "Новый текст:",
    note.querySelector("p").textContent,
  );
  if (newTitle && newContent) {
    fetch(`http://localhost:8080/notes/${id}`, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ title: newTitle, content: newContent }),
    }).then(fetchNotes);
  }
}

// Инициализация
fetchNotes();
