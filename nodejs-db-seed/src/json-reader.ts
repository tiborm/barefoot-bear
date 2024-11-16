import { promises as fs } from 'fs';
import path from "path";

async function readSingleFile(filePath: string) {
    try {
        return JSON.parse(await fs.readFile(path.resolve(filePath), 'utf8'));        
    } catch (error) {
        console.error('Error reading file:', error);
    }
}

async function readFilesInDirectory(directoryPath: string) {
    const results = [];

    try {
        const files = await fs.readdir(directoryPath);

        for (const file of files) {
            const filePath = path.join(directoryPath, file);
            const data = await fs.readFile(filePath, 'utf8');
            results.push(JSON.parse(data));
        }
    } catch (err) {
        console.error('Error reading directory or files:', err);
    }

    return results;
}

export {
    readSingleFile,
    readFilesInDirectory,
};
