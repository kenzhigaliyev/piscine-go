curl -s https://01.alem.school/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"kenzhigaliyev\"}}){id}}"}'| cut -d "}" -f1 | cut -d ":" -f4
