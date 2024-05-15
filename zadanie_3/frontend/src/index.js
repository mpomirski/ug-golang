import React from 'react';
import ReactDOM from 'react-dom';
import axios from 'axios';
import CRUDTable,
{
  Fields,
  Field,
  CreateForm,
  UpdateForm,
  DeleteForm,
} from 'react-crud-table';

// Component's Base CSS
import './index.css';

const DescriptionRenderer = ({ field }) => <textarea {...field} />;

const service = {
  fetchItems: (payload) => {
    let result = Promise.resolve(axios.get('http://localhost:3000/attacks').then(response => response.data).then(data => {
      return Object.values(data);
    }));
    return Promise.resolve(result);
  },
  create: (attack) => {
    let result = axios.post(
      'http://localhost:3000/attacks', 
      {
        id: attack.id,
        date: attack.date,
        country: attack.country,
        activity: attack.activity,
        name: attack.name,
        sex: attack.sex,
        fatal_y_n: attack.fatal_y_n
      }
    ).then(response => response.data).then(data => {
      return Object.values(data);
    });
    return Promise.resolve(result);
    // count += 1;
    // attacks.push({
    //   ...attack,
    //   id: count,
    // });
    // return Promise.resolve(attack);
  },
  update: (data) => {
    // const attack = attacks.find(t => t.id === data.id);
    // attack.title = data.title;
    // attack.description = data.description;
    // return Promise.resolve(attack);
  },
  delete: (data) => {
    let result = axios.delete('http://localhost:3000/attacks/' + data.id).then(response => response.data).then(data => {
      return Object.values(data);
    });
    return Promise.resolve(result);
  },
};

const styles = {
  container: { margin: 'auto', width: 'fit-content' },
};

const Example = () => (
  <div style={styles.container}>
    <CRUDTable
      caption="Attacks"
      fetchItems={payload => service.fetchItems(payload)}
    >
      <Fields>
      <Field
          name="id"
          label="id"
          hideInCreateForm
          readOnly
        />
        <Field
          name="date"
          label="date"
          type="date"
        />
        <Field
          name="country"
          label="country"
          placeholder="country"
        />
        <Field
          name="activity"
          label="activity"
          render={DescriptionRenderer}
        />
        <Field
          name="name"
          label="name"
          render={DescriptionRenderer}
        />
        <Field
          name="sex"
          label="sex"
          render={DescriptionRenderer}
        />
        <Field
          name="fatal_y_n"
          label="fatal_y_n"
          render={DescriptionRenderer}
        />
      </Fields>
      <CreateForm
        title="Attack Creation"
        message="Create a new attack!"
        trigger="Create Attack"
        onSubmit={attack => service.create(attack)}
        submitText="Create"
        // validate={(values) => {
        //   const errors = {};
        //   if (!values.title) {
        //     errors.title = 'Please, provide attack\'s title';
        //   }

        //   if (!values.description) {
        //     errors.description = 'Please, provide attack\'s description';
        //   }

        //   return errors;
        // }}
      />

      <DeleteForm
        title="Attack Delete Process"
        message="Are you sure you want to delete the attack?"
        trigger="Delete"
        onSubmit={attack => service.delete(attack)}
        submitText="Delete"
        // validate={(values) => {
        //   const errors = {};
        //   if (!values.id) {
        //     errors.id = 'Please, provide id';
        //   }
        //   return errors;
        // }}
      />
    </CRUDTable>
  </div>
);

Example.propTypes = {};

ReactDOM.createRoot(document.getElementById('root')).render(
  <Example />
);