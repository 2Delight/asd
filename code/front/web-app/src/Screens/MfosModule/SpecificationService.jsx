import axios from "axios";

export class SpecificationService {
    constructor(baseUrl) {
        this.baseUrl = baseUrl;
        this.axiosInstance = axios.create({
            baseURL: baseUrl,
            headers: {
                "Content-Type": "application/json",
            },
        });
    }

    // Метод для получения спецификации по ID
    async getSpecificationById(id, isMock) {
        if (isMock) {
            console.log(`Mocked fetch for specification with ID: ${id}`);
        return new Promise((resolve) => {
            setTimeout(() => resolve(this.mapToModel(mockSpecification)), 1000);
        });
        } else {
            try {
                const response = await this.axiosInstance.get(`/gateway/specifications/${id}`);
                return this.mapToModel(response.data);
            } catch (error) {
                console.error("Ошибка при получении спецификации:", error);
                throw error;
            }
        }
    }

    async updateSpecification(id, specificationContent) {
        try {
            const response = await this.axiosInstance.put(`/gateway/specifications/${id}`, {
                id,
                specificationContent,
            });
            return response.data;
        } catch (error) {
            console.error("Ошибка при обновлении спецификации:", error);
            throw error;
        }
    }

    async validateSpecification(id, specificationContent, isMock = false) {
        if (isMock) {
            console.log(`Mocked validation for specification with ID: ${id}`);
            return new Promise((resolve) => {
                // setTimeout(() => resolve(mockValidationResult), 1000);
                setTimeout(() => resolve(mockValidationResult), 1000);
            });
        } else {
            try {
                const response = await this.axiosInstance.post(`/gateway/specifications/${id}/validate`, {
                    id,
                    specificationContent,
                });
                return response.data;
            } catch (error) {
                console.error("Ошибка при валидации спецификации:", error);
                throw error;
            }
        }
    }

    // Метод для маппинга ответа на модель
    mapToModel(data) {
        return {
            id: data.id,
            name: data.name,
            content: data.content,
            gitPath: data.gitPath,
            createdAt: new Date(data.createdAt),
            updatedAt: new Date(data.updatedAt),
        };
    }
}

// Замоканные данные для тестирования
const mockSpecification = {
    id: 1,
    name: "Specification Example",
    content: `id: 1\nname: SampleSpecification\nfields:\n  - name: field1\n    type: string\n    required: true\n  - name: field2\n    type: integer\n    required: false\n  - name: field3\n    type: boolean\n    required: true\nmetadata:\n  createdBy: John Doe\n  createdAt: 2024-01-01T12:00:00Z\n  updatedAt: 2024-01-02T12:00:00Z\n  tags:\n    - spec\n    - example\n    - mock\nrelationships:\n  - target: OtherSpec\n    type: one-to-one\n  - target: AnotherSpec\n    type: one-to-many\nexamples:\n  example1:\n    field1: "exampleString"\n    field2: 42\n    field3: true\n  example2:\n    field1: "anotherString"\n    field2: null\n    field3: false\nspecification:\n  description: "This is a sample specification for testing purposes."\n  details:\n    - detailName: Detail1\n      detailValue: "Value1"\n    - detailName: Detail2\n      detailValue: "Value2"\n    - detailName: Detail3\n      detailValue: "Value3"\n  notes:\n    - "This specification is a mock example."\n    - "Feel free to extend this example for real use cases."\n  compliance:\n    version: "1.0"\n    status: "Approved"\n  lifecycle:\n    createdAt: "2024-01-01T12:00:00Z"\n    updatedAt: "2024-01-02T12:00:00Z`,
    gitPath: "/path/to/git",
    createdAt: "2024-01-01T12:00:00Z",
    updatedAt: "2024-01-02T12:00:00Z",
};

const mockValidationResult = {
    isValid: false,
    errors: [
        {
            code: "INVALID_FIELD",
            message: "Field 'name' is missing or invalid",
        },
        {
            code: "MISSING_FIELD",
            message: "Field 'createdBy' is required but not provided",
        },
    ],
    hints: [
        {
            message: "Ensure all required fields are present",
        },
        {
            message: "Check the format of 'name' field",
        },
    ],
};

const mockValidationSuccessResult = {
    isValid: true,
    errors: [],
    hints: [],
};