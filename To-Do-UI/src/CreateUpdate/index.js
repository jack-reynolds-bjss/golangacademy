import axios from 'axios';
import { useEffect, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { Button, Input, InputLabel, MenuItem, Select } from '@mui/material';
import './index.css';

const axiosConfig = {
  headers: {
    'Content-Type': 'application/json'
  }
}

const defaultToDo = {
    item: "",
    status: "NOTSTARTED"
}

const CreateUpdate = ({ edit = false }) => {
  const navigate = useNavigate();
  const { id } = useParams();
  const [todo, setTodo] = useState(defaultToDo);

  useEffect(() => {
    if (id) {
      axios.get('http://localhost:8090/todoItems', { params: { id } }).then(response => {
        if (response.data) {
          setTodo(response.data)
        } 
      });
    }
  }, [id])

  const onChangeItem = (val) => {
      setTodo({
          ...todo,
          item: val
      })
  }

  const onChangeStatus = (val) => {
      setTodo({
          ...todo,
          status: val
      })
  }

  const onSave = () => {
    if (edit) {
      axios.post('http://localhost:8090/todoItems', todo, axiosConfig).then(() => {
        navigate('/')
      });
    } else {
      axios.put('http://localhost:8090/todoItems', todo, axiosConfig).then(() => {
        navigate('/')
      });
    }
  }

  return (
    <div className='create-edit-todo'>
      <h2>{`${edit ? "Edit" : "Create"} ToDo`}</h2>
      <div className='todo-form'>
        <div className='item-container'>
          <InputLabel>Item</InputLabel>
          <Input name="item" title='Item' sx={{ width: '100%' }} onChange={(e) => onChangeItem(e.target.value)} value={todo.item} />
        </div>
        <div className='status-container'>
          <InputLabel>Status</InputLabel>
          <Select
            name="status"
            title="Status"
            label="Status"
            placeholder="Status"
            value={todo.status}
            onChange={(e) => onChangeStatus(e.target.value)}
            sx={{ width: '100%' }}
          >
            <MenuItem value="NOTSTARTED">Not Started</MenuItem>
            <MenuItem value="INPROGRESS">In Progress</MenuItem>
            <MenuItem value="COMPLETED">Completed</MenuItem>
          </Select>
        </div>
        </div>
      <div className='save-container'><Button onClick={onSave}>Save</Button></div>
    </div>
  );
}

export default CreateUpdate;
