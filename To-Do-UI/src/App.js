import axios from 'axios';
import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { Button, IconButton } from '@mui/material';
import { Delete as DeleteIcon, Close as CloseIcon, Pending as PendingIcon, CheckCircle as CheckCircleIcon, QuestionMark as QuestionMarkIcon} from '@mui/icons-material';

import './App.css';

const getStatusIcon = (status) => {
  switch (status) {
    case "NOTSTARTED":
      return <CloseIcon sx={{ color: 'red' }} />
    case "INPROGRESS":
      return <PendingIcon />
    case "COMPLETED":
      return <CheckCircleIcon sx={{ color: 'green' }} />
    default:
      return <QuestionMarkIcon />
  }
}

const Todo = ({ todo, onClickTodo, onClickDelete }) => {
  return <div key={todo.id} onClick={onClickTodo} className='todo'>
      <div className='status'>{getStatusIcon(todo.status)}</div>
      <div className='item'>{todo.item}</div>
      <div className='delete-container'><IconButton onClick={onClickDelete} className='delete'><DeleteIcon /></IconButton></div>
    </div>
};

const App = () => {
  const navigate = useNavigate();
  const [items, setItems] = useState([]);

  const onClickAdd = () => {
    navigate('/create')
  }

  const onClickTodo = (e, id) => {
    e.preventDefault()
    navigate(`/edit/${id}`);
  }

  const onClickDeleteTodo = (e, id) => {
    e.preventDefault()
    axios.delete('http://localhost:8090/todoItems', { params: { id } }).then(() => {
      navigate('/');
    });
  }

  useEffect(() => {
    axios.get('http://localhost:8090/todoItems').then(response => {      
      if (response.data?.length) {
        setItems(response.data)
      } 
    });
  }, []);

  return (
    <div className='app'>
      <div className='action-bar'><Button onClick={onClickAdd}>Add</Button></div>
      {items.map(todo => <Todo todo={todo} onClickTodo={(e) => onClickTodo(e, todo.id)} onClickDelete={(e) => onClickDeleteTodo(e, todo.id)} />)}
    </div>
  );
}

export default App;
