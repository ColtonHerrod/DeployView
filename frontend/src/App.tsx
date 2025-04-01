import React, { useState } from 'react';
import axios from 'axios';

const App: React.FC = () => {
    const [accountId, setAccountId] = useState('');
    const [roleArn, setRoleArn] = useState('');
    const [applicationName, setApplicationName] = useState('');
    const [result, setResult] = useState<any>(null);

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        try {
            const response = await axios.post(`${process.env.REACT_APP_API_URL}/codedeploy`, {
                accountId,
                roleArn,
                applicationName,
            });
            setResult(response.data);
        } catch (error) {
            console.error(error);
            alert('Error fetching CodeDeploy information.');
        }
    };

    return (
        <div>
            <h1>AWS CodeDeploy Viewer</h1>
            <form onSubmit={handleSubmit}>
                <div>
                    <label>Account ID:</label>
                    <input
                        type="text"
                        value={accountId}
                        onChange={(e) => setAccountId(e.target.value)}
                        required
                    />
                </div>
                <div>
                    <label>Role ARN:</label>
                    <input
                        type="text"
                        value={roleArn}
                        onChange={(e) => setRoleArn(e.target.value)}
                        required
                    />
                </div>
                <div>
                    <label>Application Name:</label>
                    <input
                        type="text"
                        value={applicationName}
                        onChange={(e) => setApplicationName(e.target.value)}
                        required
                    />
                </div>
                <button type="submit">Fetch CodeDeploy Info</button>
            </form>
            {result && (
                <div>
                    <h2>Application Details</h2>
                    <pre>{JSON.stringify(result, null, 2)}</pre>
                </div>
            )}
        </div>
    );
};

export default App;
