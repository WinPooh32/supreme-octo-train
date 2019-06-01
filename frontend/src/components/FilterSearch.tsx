import * as React from "react";
import { Input } from 'semantic-ui-react'

export class FilterSearch extends React.Component<{}, {}> {
    render() {
        return <Input fluid icon='search' placeholder='Фильтр...' />;
    }
}