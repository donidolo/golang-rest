import { db } from './firebase';
import { Note, NewNote } from '@/store/models';
import axios from 'axios';

const noteCollection = db.collection('notes');

export const createNote = async (note: NewNote): Promise<Note> => {
  const newDoc = await noteCollection.add(note);
  return {
    id: newDoc.id,
    ...note
  };
};

export const getNotes = async (): Promise<Note[]> => {
  /*
  fetch('http://localhost:8081/notes?origin='+location.origin,{ mode: 'no-cors' })
    .then(response => {
      debugger
      response.json()
    })
    .then(data => {
      debugger
      console.log(data)
    })
    .catch(err => {
      debugger
    });
    */
  // const url = 'http://localhost:8081/notes?origin=' + location.origin
  // const headers = {
  //   'Access-Control-Allow-Origin': '*',
  // }

  const apiClient = axios.create({
    baseURL: "http://localhost:8081",
    withCredentials: false,
    headers: {
      Accept: 'application/json',
    },
  })

  apiClient.interceptors.request.use(c => {
    c.headers.token="62532165376"
    c.params = {
      ...c.params,
      origin: location.origin,
    }

    return c
  })

  const resp = await apiClient.get("/notes")
  const notes = resp.data
  return notes
  // const querySnapshot = await noteCollection.get();
  // const temp = querySnapshot.docs.map(doc => {
  //   const { title, content, color } = doc.data();
  //   return {
  //     id: doc.id,
  //     title,
  //     content,
  //     color
  //   };
  // });
  // console.log(temp);
  // debugger
  // return temp
};

export const updateNote = async (note: Note): Promise<Note> => {
  await noteCollection.doc(note.id).update({
    title: note.title,
    content: note.content,
    color: note.color
  });
  return note;
};

export const deleteNote = async (id: string): Promise<string> => {
  await noteCollection.doc(id).delete();
  return id;
};
